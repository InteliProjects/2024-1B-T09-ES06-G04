import React, { useEffect } from 'react';
import { View, ScrollView, TouchableOpacity, Alert } from 'react-native';
import TextArea from '../../components/TextArea';
import Icon from 'react-native-vector-icons/Ionicons';
import CategoryButton from '../../components/CategoryButton';
import Title from '../../components/Title';
import styles from './styles';
import { decodeJWT } from '../../services/decodeJWT';
import { useState } from 'react';
import { getToken, projectsApi, registerProject } from '../../services/api';


// This screen is responsible for adding a new project
export default function AddProjectScreen({navigation}) {
  const [userId, setUserId] = useState(null);
  const [name, setName] = useState('');
  const [macroSector, setMacroSector] = useState('');
  const [microSector, setMicroSector] = useState('');
  const [imageLink, setImageLink] = useState('');
  const [description, setDescription] = useState('');

  // Get the user ID from the token
  useEffect(() => {
    const loadData = async () => {
      const token = await getToken();
      if (token) {
        const decodedToken = decodeJWT(token);
        setUserId(decodedToken.id);
      }
    };

    loadData();
  }, []);

  // Function to handle the addition of a new project
  const handleAddProject = async () => {
    console.log('Sending project data:', { name, description, macroSector, microSector, imageLink, userId });
  
    if (!userId) {
      console.error('User ID is not set.');
      Alert.alert('Erro', 'Não foi possível identificar o usuário.');
      return;
    }
  
    try {
      const response = await registerProject(name, description, macroSector, microSector, imageLink, parseInt(userId));
      console.log('Project registered:', response);
      Alert.alert('Sucesso!', 'Projeto cadastrado com sucesso!');
    } catch (error) {
      console.error('Failed to register project:', error);
      Alert.alert('Erro ao cadastrar projeto', 'Tente novamente mais tarde');
    }
  };
  
  return (
    <View style={styles.container}>
      <View style={styles.header}> 
        <Title 
          title="Novo Projeto"
        />
        <TouchableOpacity onPress={() => navigation.goBack()} style={styles.closeButton}>
          <Icon name="close" size={30} color="#000" />
        </TouchableOpacity>
      </View>
      <ScrollView contentContainerStyle={styles.scrollContent}>
        <View style={styles.header}>
        </View>
        <TextArea 
          title="Nome do Projeto" 
          minHeight={50} 
          value={name}
          onChangeText={setName}
        />
        <TextArea 
          title="Macro Setor" 
          minHeight={50} 
          value={macroSector}
          onChangeText={setMacroSector}
        />
        <TextArea 
          title="Micro Setor" 
          minHeight={50} 
          value={microSector}
          onChangeText={setMicroSector}
        />
        <TextArea 
          title="Link de Imagem de demonstração" 
          minHeight={50} 
          value={imageLink}
          onChangeText={setImageLink}
        />
        <TextArea 
          title="Descrição" 
          minHeight={250} 
          value={description}
          onChangeText={setDescription}
        />
      </ScrollView>
      <View style={styles.buttonContainer}>
        <CategoryButton 
          title="Adicionar"
          onPress={handleAddProject}
          backgroundColor="#B6E99E"
          color="#000"
        />
      </View>
    </View>
  );
}
