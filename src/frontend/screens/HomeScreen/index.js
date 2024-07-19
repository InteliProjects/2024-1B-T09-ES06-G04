import { View, ScrollView, Alert } from 'react-native';
import React, { useEffect, useState } from 'react';
import Card from '../../components/Card';
import styles from './styles';
import CategoryButton from '../../components/CategoryButton';
import { projectsApi } from '../../services/api';
import { useFocusEffect } from '@react-navigation/native';  
import { useCallback } from 'react';

// This screen is responsible for adding a home page
export default function HomeScreen({ navigation }) {
  const [projects, setProjects] = useState([]);
  const [filteredProjects, setFilteredProjects] = useState([]);


  // This function fetches the projects from the API
  const fetchProjects = async () => {
    try {
      const response = await projectsApi.get('/projects');
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

  // This function filters the projects by category
  const handleCategoryPress = (category) => {
    const filtered = projects.filter(project => project.macro_setor === category);
    setFilteredProjects(filtered);
  };

  // This function resets the filter and shows all projects
  const resetFilter = () => {
    setFilteredProjects(projects);
  };

  return (
    <ScrollView style={styles.container}>
      <View style={styles['container-scroll']}>
        <ScrollView horizontal={true} showsHorizontalScrollIndicator={false}>
          <CategoryButton
            title={'Todos os Projetos'}
            onPress={resetFilter}
            backgroundColor={'#000'}
            color={'#fff'}
          />
          <CategoryButton
            title={'Saúde e bem-estar'}
            onPress={() => handleCategoryPress('Saúde e Bem-estar')}
            backgroundColor={'#ECC087'}
          />
          <CategoryButton
            title={'DE&I'}
            onPress={() => handleCategoryPress('Serviços Sociais')}
            backgroundColor={'#DF9D4B'}
            color={'#fff'}
          />
          <CategoryButton
            title={'Redução do Impacto Ambiental'}
            onPress={() => handleCategoryPress('Meio Ambiente')}
            backgroundColor={'#4E5F3D'}
            color={'#fff'}
          />
          <CategoryButton
            title={'Integridade e Práticas Éticas'}
            onPress={() => handleCategoryPress('Ciência e Pesquisa')}
            backgroundColor={'#81859E'}
            color={'#fff'}
          />
          <CategoryButton
            title={'Produtividade e Competitividade'}
            onPress={() => handleCategoryPress('Educação')}
            backgroundColor={'#454861'}
            color={'#fff'}
          />
        </ScrollView>
      </View>
      {filteredProjects.map((project) => (
        <Card
          key={project.id}
          projectId={project.id}
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
