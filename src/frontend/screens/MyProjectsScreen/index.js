import React, { useEffect, useState } from 'react';
import { View, ScrollView, Alert, Modal, Text, TouchableOpacity } from 'react-native';
import Title from '../../components/Title'; 
import Card from '../../components/Card'; 
import styles from './styles';
import { getProjectsByUser, deleteProject, projectsApi  } from '../../services/api';
import Input from '../../components/Input';
import TextArea from '../../components/TextArea';
import IconClose from 'react-native-vector-icons/AntDesign';
import Button from '../../components/Button';
import { updateProject, getToken } from '../../services/api';
import { decodeJWT } from '../../services/decodeJWT';

// This screen is responsible for displaying the user's projects
export default function MyProjectsScreen({ navigation }) {
  const [projects, setProjects] = useState([]);
  const [modalVisible, setModalVisible] = useState(false);
  const [currentProject, setCurrentProject] = useState(null);
  const [projectName, setProjectName] = useState('');
  const [macroSector, setMacroSector] = useState('');
  const [microSector, setMicroSector] = useState('');
  const [imageLink, setImageLink] = useState('');
  const [description, setDescription] = useState('');
  const [userId, setUserId] = useState(0);
  const [projectId, setProjectId] = useState(0);

  // Get the user ID from the token
  useEffect(() => {
    const loadData = async () => {
      const token = await getToken();
      if (token) {
        const decodedToken = decodeJWT(token);
        setUserId(parseInt(decodedToken.id));
        console.log('User ID:', decodedToken.id);
      }
    };
    loadData();
  }, []);

  // Get the user's projects
  useEffect(() => {
    const fetchProjects = async () => {
      try {
        const userProjects = await getProjectsByUser();
        setProjects(userProjects);
      } catch (error) {
        Alert.alert('Erro', 'Não foi possível obter os projetos do usuário.');
      }
    };

    fetchProjects();
  }, []);

  // Set the project data to the modal fields
  useEffect(() => {
    if (currentProject) {
      setProjectId(currentProject.id);
      setProjectName(currentProject.name);
      setMacroSector(currentProject.macro_setor);
      setMicroSector(currentProject.micro_setor);
      setImageLink(currentProject.image_link);
      setDescription(currentProject.description);
    }
  }, [currentProject]);

  // Function to handle the deletion of a project
  const handleDeleteProject = async (projectId) => {
    try {
        await deleteProject(projectId);
        Alert.alert('Sucesso', 'Projeto deletado com sucesso');
        setProjects(projects.filter(project => project.id !== projectId));
    } catch (error) {
        Alert.alert('Erro', 'Falha ao deletar o projeto');
    }
  };

  // Function to handle the update of a project
  const handleUpdateProject = async () => {
    try{
      const updatedData = {
        ...currentProject,
        id: projectId,
        name: projectName,
        description: description,
        macro_setor: macroSector,
        micro_setor: microSector,
        image_link: imageLink,
        user_id: userId
      };

      const response = await projectsApi.put(`/projects/${projectId}`, updatedData);
      console.log('Projeto atualizado com sucesso:', response.data);
      Alert.alert('Sucesso', 'Projeto atualizado com sucesso.');
      } catch (error) {
        console.error('Falha ao atualizar o projeto:', error);
    }
  };

  // Function to open the modal with the project data
  const openModal = (project) => {
    setCurrentProject(project);
    setModalVisible(true);
  };

  // Function to close the modal
  const handleClose = () => {
    setModalVisible(false);
  };

  return (
    <>
    <ScrollView style={styles.container}>
      <Title title="Seus projetos" />
      {projects.map((project) => (
        <Card 
          key={project.id}
          title={project.name}
          description={project.description}
          MyProject={true}
          navigation={navigation}
          handleDeleteProject={() => handleDeleteProject(project.id)}
          handleEditProject={() => openModal(project)}
        />
      ))}
      </ScrollView>

      <Modal
        animationType="slide"
        transparent={false}
        visible={modalVisible}
        onRequestClose={handleClose}
        contentContainerStyle={{ alignItems: 'center' }}
      >
        <View style={styles.modalView}>
          <TouchableOpacity style={styles.closeButton} onPress={handleClose}>
            <Text style={styles.closeButtonText}>
              <IconClose name='close' size={30} color='#000' />
            </Text>
          </TouchableOpacity>
          <Input 
            label="Nome do Projeto"
            value={projectName}
            style={{ marginBottom: 20 }}
            onChangeText={setProjectName}
          />
          <Input 
            label="Macro Setor"
            value={macroSector}
            style={{ marginBottom: 20 }}
            onChangeText={setMacroSector}
          />
          <Input 
            label="Micro Setor"
            value={microSector}
            style={{ marginBottom: 20 }}
            onChangeText={setMicroSector}
          />
          <Input 
            label="Link da Imagem de demonstração"
            value={imageLink}
            style={{ marginBottom: 20 }}
            onChangeText={setImageLink}
          />
          <TextArea 
            title="Descrição"
            value={description}
            style={styles.descriptionInput}
            minHeight={100}
            onChangeText={setDescription}
          />
        <Button label="Salvar" onPress={handleUpdateProject} style={styles.saveButton}/>
        </View>
      </Modal>
      </>
  );
}
