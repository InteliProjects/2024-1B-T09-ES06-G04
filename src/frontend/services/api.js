import axios from "axios";
import AsyncStorage from '@react-native-async-storage/async-storage';

// Axios instance for the authentication service
const authApi = axios.create({
    baseURL: 'http://100.29.135.172:8082/api/v1'
});

// Axios instnce for the projects service
const projectsApi = axios.create({
    baseURL: 'http://100.29.135.172:8082/api/v1'
});

// Axios instance for the avatars service
const avatarsApi = axios.create({
    baseURL: 'http://100.29.135.172:8082'
});

// Axios instance for the connections service
const connectionsApi = axios.create({
    baseURL: 'http://100.29.135.172:8082/api/v1'
});

// Function to store the token in AsyncStorage
const storeToken = async (token) => {
    try {
        await AsyncStorage.setItem('authToken', token);
    } catch (e) {
        console.error('Saving token failed', e);
    }
};

// Function to load the token from AsyncStorage
const getToken = async () => {
    try {
        return await AsyncStorage.getItem('authToken');
    } catch (e) {
        console.error('Loading token failed', e);
        return null;
    }
};

// Function to remove the token from AsyncStorage
const removeToken = async () => {
    try {
        await AsyncStorage.removeItem('authToken');
        console.log('Token removed successfully');
    } catch (e) {
        console.error('Failed to remove token', e);
    }
};

// Interceptors to add the authentication token to all requests
const setupInterceptors = (api) => {
    api.interceptors.request.use(
        async (config) => {
            const token = await getToken();
            if (token) {
                config.headers['Authorization'] = `Bearer ${token}`;
            }
            return config;
        },
        error => {
            return Promise.reject(error);
        }
    );
};

// Configure interceptors for both instances
setupInterceptors(authApi);
setupInterceptors(projectsApi);

// Function to log in using the authentication service
const login = async (email, password) => {
    try {
        const response = await authApi.post('/login', { email, password });
        const { token } = response.data;
        await storeToken(token);
        return token; 
    } catch (error) {
        throw error;
    }
};

// Function to register a new user using the authentication service
const register = async (name, email, password, companyName, office, linkedinLink, interest) => {
    try {
        const response = await authApi.post('/register', {
            name,
            email,
            password,
            company_name: companyName, 
            office,
            linkedin_link: linkedinLink, 
            interest
        });
        console.log('Register response:', response.data);
        return response.data;
    } catch (error) {
        console.error('Registration error:', error);
        throw error;
    }
};

// Function to register a new project using the projects service
const registerProject = async (name, description, macroSector, microSector, imageLink, userId) => {
    try {
        const response = await projectsApi.post('/projects', {
            name,
            description,
            macro_setor: macroSector,
            micro_setor: microSector,
            image_link: imageLink,
            user_id: userId
        });
        console.log('Project registered successfully:', response.data);
        return response.data;
    } catch (error) {
        console.error('Failed to register project:', error);
        throw error;
    }
};

// Function to delete a project using the projects service
const deleteProject = async (projectId) => {
    try {
        const response = await projectsApi.delete(`/projects/${projectId}`);
        console.log('Project deleted successfully:', response.data);
    } catch (error) {
        console.error('Failed to delete project:', error);
        throw error;
    }
};

// Function to update the project by id project using the projects service
const updateProject = async (projectId, name, description, macroSector, microSector, imageLink, userId) => {
    try {
        const response = await projectsApi.put(`/projects/${projectId}`, {
            name,
            description,
            macro_setor: macroSector,
            micro_setor: microSector,
            image_link: imageLink,
            user_id: userId
        });
        console.log('Project updated successfully:', response.data);
    } catch (error) {
        console.error('Failed to update project:', error);
        throw error;
    }
};


// Function to update the user interests using the authentication service
const updateUserInterests = async (userId, interests) => {
    try {
        const response = await authApi.put(`/users/${userId}`, { interest: interests.join(', ') });
        console.log('Interests updated successfully:', response.data);
    } catch (error) {
        console.error('Failed to update interests:', error);
    }
}

// Function to get the avatars from the avatars service
const getAvatars = async () => {
    try {
        const response = await avatarsApi.get('/avatars');
        return response.data;
    } catch (error) {
        console.error('Failed to get avatars:', error);
        return [];
    }
};

// Function to get projects by user
const getProjectsByUser = async () => {
    try {
        const response = await projectsApi.get('/projects/user');
        return response.data;
    } catch (error) {
        console.error('Failed to get user projects:', error);
        throw error;
    }
};

export { authApi, projectsApi, connectionsApi, login, storeToken, getToken, register, updateUserInterests, getAvatars, getProjectsByUser, removeToken, registerProject, deleteProject, updateProject};
