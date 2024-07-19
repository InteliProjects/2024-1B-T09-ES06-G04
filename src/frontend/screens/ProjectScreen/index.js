import React, { useState, useCallback } from 'react';
import { View, Text, Image, TouchableOpacity, ScrollView, Alert } from 'react-native';
import { useFocusEffect } from '@react-navigation/native';
import styles from './styles';
import Icon from 'react-native-vector-icons/Ionicons';
import IconRecommend from 'react-native-vector-icons/AntDesign';
import { projectsApi } from '../../services/api';
import Title from '../../components/Title';
import IconGreen from '../../components/IconGreen';

//This screen is responsible for presenting information about a project
export default function ProjectScreen({ navigation, route }) {
  const [project, setProject] = useState(null);
  const [connections, setConnections] = useState([]);

  // Function to fetch project data by ID
  const fetchProjectById = async (id) => {
    try {
      console.log('Fetching project with ID:', id);
      const response = await projectsApi.get(`/projects/${id}`);
      const projectData = response.data;

      // Fetch user data related to the project
      const userResponse = await projectsApi.get(`/users/${projectData.user_id}`);
      const userData = userResponse.data;

      // Combine project data with user data
      const projectWithUser = {
        ...projectData,
        user_name: userData.name,
        user_company: userData.company_name,
      };

      // Set the project and connections state
      setProject(projectWithUser);
      setConnections(projectData.connections);

    } catch (err) {
      if (err.response) {
        Alert.alert('Erro na resposta', 'Não foi possível obter o projeto.');
      } else {
        Alert.alert('Erro de conexão', 'Verifique sua conexão de internet.');
      }
    }
  };

  // Hook to fetch project data when the screen is focused
  useFocusEffect(
    useCallback(() => {
      const id = route.params?.id;
      console.log('Project ID from route params:', id);
      if (id) {
        fetchProjectById(id);
      }
    }, [route.params?.id])
  );

  // If project data is not yet available, display a loader
  if (!project) {
    return (
      <View style={styles.loaderContainer}>
        <Text>Projeto não encontrado.</Text>
      </View>
    );
  }

  // Render the project screen with details
  return (
    <View style={styles.outerContainer}>
      <ScrollView contentContainerStyle={styles.container}>
        <View style={styles.header}>
          <Title title={project.name} subtitle={project.macro_setor} />
          <TouchableOpacity onPress={() => navigation.goBack()} style={styles.backButton}>
            <Icon name="arrow-back" size={24} color="#000" />
          </TouchableOpacity>
        </View>

        <View style={styles.coverPhotoContainer}>
          <Image source={{ uri: project.image_link || 'https://via.placeholder.com/150' }} style={styles.coverPhoto} />
        </View>

        <View style={styles.connectButtonContainer}>
          <IconGreen image={require('../../assets/icon_perfil.png')} name={project.user_name} company={project.user_company} />
          <TouchableOpacity style={styles.connectButton}>
            <IconRecommend name="like2" size={20} color="#000" style={styles.connectButtonIcon} />
            <Text style={styles.connectButtonText}>Conectar</Text>
          </TouchableOpacity>
        </View>

        <View style={styles.descriptionContainer}>
          <Text style={styles.sectionTitle}>Descrição</Text>
          <Text style={styles.description}>{project.description}</Text>
        </View>

        <View style={styles.connectionsContainer}>
          <Text style={styles.sectionTitle}>Conexões Realizadas</Text>
          <View style={styles.connection}>
            <Image
              source={require('../../assets/icon_perfil.png')}
              style={styles.connectionImage}
            />
            <Text>Usuário 1 | Projeto: Nome do Projeto</Text>
          </View>
          <View style={styles.connection}>
            <Image
              source={require('../../assets/icon_perfil.png')}
              style={styles.connectionImage}
            />
            <Text>Usuário 2 | Projeto: Nome do Projeto</Text>
          </View>
          <View style={styles.connection}>
            <Image
              source={require('../../assets/icon_perfil.png')}
              style={styles.connectionImage}
            />
            <Text>Usuário 3 | Projeto: Nome do Projeto</Text>
          </View>
        </View>
      </ScrollView>
    </View>
  );
}
