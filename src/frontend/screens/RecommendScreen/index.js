import { ScrollView } from "react-native";
import Card from "../../components/Card";
import styles from "./styles";
import Title from "../../components/Title";
import { projectsApi } from '../../services/api';
import { useEffect, useState } from "react";
import { useCallback } from "react";
import { Alert } from "react-native";


// This screen is responsible for displaying the recommended projects
export default function RecommendScreen({navigation}) {
  const [projects, setProjects] = useState([]);

  // This function fetches the projects from the API
  const fetchProjects = async () => {
    try {
      const response = await projectsApi.get('/projects');
      setProjects(response.data);
      setFilteredProjects(response.data);
    } catch (err) {
      if (err.response) {
        Alert.alert('Erro na resposta', 'Não foi possível obter os projetos.');
      } 
    }
  };
 
  // This hook fetches the projects from the API when the component is mounted
  useEffect(
    useCallback(() => {
      fetchProjects();
    }, [])
  );

  return (
    <ScrollView style={styles.container}>
      <Title 
        title="Projetos Recomendados"
      />
      {projects.map((project) => (
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