import React, { useState, useEffect } from 'react';
import { Modal, Text, ScrollView, View, TouchableOpacity, Alert } from 'react-native';
import { useNavigation } from '@react-navigation/native';
import styles from './styles';
import IconClose from 'react-native-vector-icons/AntDesign';
import IconFontAwesome from 'react-native-vector-icons/FontAwesome';
import Input from '../../components/Input';
import { register } from '../../services/api'

//This screen is responsible for registering new users
export default function RegisterScreen() {
  const [modalVisible, setModalVisible] = useState(true);
  const navigation = useNavigation();
  const [name, setName] = useState('');
  const [companyName, setCompanyName] = useState('');
  const [office, setOffice] = useState('');
  const [linkedinLink, setLinkedinLink] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  // Function to handle the registration of a new user
  const handleRegister = async () => {
    try {
      console.log('Attempting to register:', { name, email, password, companyName, office, linkedinLink });
      const response = await register(name, email, password, companyName, office, linkedinLink, null);
      console.log('Registration successful:', response);
      
      navigation.navigate('Interest', { userId: response.userID });
    } catch (error) {
      console.error('Register error:', error);
      if (error.response) {
          Alert.alert('Erro ao cadastrar', `Erro: ${error.response.data.message}`);
      } else {
          Alert.alert('Erro ao cadastrar', 'Tente novamente mais tarde');
      }
    }
  };

  // Function to close the modal and navigate to the login screen
  useEffect(() => {
    const unsubscribe = navigation.addListener('focus', () => {
      setModalVisible(true);
    });

    return unsubscribe;
  }, [navigation]);

  // Function to close the modal and navigate to the login screen
  const handleClose = () => {
    setModalVisible(false);
    navigation.navigate('Login');
  };

  return (
    <View style={styles.centeredView}>
      <Modal
        animationType="slide"
        transparent={true}
        visible={modalVisible}
        onRequestClose={handleClose}
        style={{ flex: 1 }}
      >
        <ScrollView style={styles.fullScreenModalView} contentContainerStyle={{ alignItems: 'center'}}>
          <TouchableOpacity
            style={styles.closeButton}
            onPress={handleClose}
          >
            <Text style={styles.closeButtonText}>
              <IconClose name="close" size={30} color="#000" />
            </Text>
          </TouchableOpacity>

          <View style={styles.profile}>
            <IconFontAwesome name="user-circle-o" size={100} color="#000" />
          </View>

          <View style={styles.inputContainer}>
            <Input 
              label="Nome Completo" 
              placeholder="Nome" 
              style={{width: '100%'}}
              value={name}
              onChangeText={setName}
            />
          </View>

          <View style={styles.inputContainer}>
            <Input 
              label="Nome da Empresa" 
              placeholder="Nome da empresa" 
              style={{width: '100%'}}
              value={companyName}
              onChangeText={setCompanyName}
            />
          </View>

          <View style={styles.inputContainer}>
            <Input 
              label="Cargo" 
              placeholder="Cargo" 
              style={{width: '100%'}}
              value={office}
              onChangeText={setOffice}
            />
          </View>

          <View style={styles.inputContainer}>
            <Input
              label="Perfil do LinkedIn"
              placeholder="Perfil do LinkedIn"
              style={{width: '100%'}}
              value={linkedinLink}
              onChangeText={setLinkedinLink}
            />
          </View>

          <View style={styles.inputContainer}>
            <Input 
              label="Email" 
              placeholder="Email" 
              style={{width: '100%'}}
              value={email}
              onChangeText={setEmail}
            />
          </View>

          <View style={styles.inputContainer}>
            <Input
              label="Senha"
              placeholder="Senha"
              secureTextEntry={true}
              style={{width: '100%'}}
              value={password}
              onChangeText={setPassword}
            />
          </View>

          <TouchableOpacity style={styles.button} onPress={handleRegister}>
            <Text style={styles.textStyle}>Cadastrar</Text>
          </TouchableOpacity>
        </ScrollView>
      </Modal>
    </View>
  );
}
