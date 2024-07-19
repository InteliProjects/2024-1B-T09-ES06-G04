import { View, ScrollView, Alert } from 'react-native';
import React, { useEffect, useState } from 'react';
import CardTimeline from '../../components/CardTimeline';
import styles from './styles';
import CategoryButton from '../../components/CategoryButton';
import { projectsApi } from '../../services/api';
import { useFocusEffect } from '@react-navigation/native';
import { useCallback } from 'react';

// This screen is responsible for adding a home page
export default function TimelineScreen({ navigation }) {
    const [projects, setProjects] = useState([]);
    const [filteredProjects, setFilteredProjects] = useState([]);


    // This function fetches the projects from the API
    const fetchProjects = async () => {
        try {
            const response = await projectsApi.get('/connections/true');
            setProjects(response.data);
            setFilteredProjects(response.data);
        } catch (err) {
            if (err.response) {
                Alert.alert('Erro na resposta', 'Não foi possível obter os projetos.');
            } else {
                Alert.alert('Erro de conexão', 'Verifique sua conexão de internet.');
            }
        }
    };

    // This hook fetches the projects from the API when the component is mounted
    useFocusEffect(
        useCallback(() => {
            fetchProjects();
        }, [])
    );

    return (
        <ScrollView style={styles.container}>
            <View style={styles['container-scroll']}>
                {filteredProjects.map((project) => (
                    <CardTimeline
                        key={project.id}
                        projectId={project.id}
                        title={project.project_name}
                        value={project.feedback}
                        nameUser={project.project_user_name}
                        nameConnection={project.user_name}
                        imageUser={require('../../assets/icon_perfil.png')}
                        navigation={navigation}
                    />
                ))}
            </View>
        </ScrollView>
    );
}
