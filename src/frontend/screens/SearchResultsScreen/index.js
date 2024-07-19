import React, { useEffect, useState } from 'react';
import { View, ScrollView, Alert } from 'react-native';
import Card from '../../components/Card';
import { projectsApi } from '../../services/api';

//This screen is responsible for showing search results
export default function SearchResultsScreen({ route, navigation }) {
  const { searchText } = route.params;
  const [projects, setProjects] = useState([]);
  const [filteredProjects, setFilteredProjects] = useState([]);

  // This hook fetches the projects from the API when the component is mounted
  useEffect(() => {
    fetchProjects();
  }, []);

  // This function fetches the projects from the API
  const fetchProjects = async () => {
    try {
      const response = await projectsApi.get('/projects');
      setProjects(response.data);
      filterProjects(searchText, response.data);
    } catch (err) {
      Alert.alert('Erro', 'Não foi possível obter os projetos.');
    }
  };
  
  // This function filters the projects by name
  const filterProjects = (searchText, projects) => {
    const filtered = projects.filter(project =>
      project.name.toLowerCase().includes(searchText.toLowerCase())
    );
    setFilteredProjects(filtered);
  };
  
  return (
    <ScrollView>
      {filteredProjects.map((project) => (
        <Card
          key={project.id}
          title={project.name}
          description={project.description}
          nameUser={project.user_name}
          companyUser={project.user_company}
          imageUser={require('../../assets/icon_perfil.png')}
          navigation={navigation}
        />
      ))}
    </ScrollView>
  );
}
