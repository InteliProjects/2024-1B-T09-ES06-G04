import React, { useState, useEffect } from 'react';
import {
  Modal,
  ScrollView,
  View,
  Text,
  TouchableOpacity,
  Alert,
  ActivityIndicator
} from 'react-native';
import styles from './styles';
import IconClose from 'react-native-vector-icons/AntDesign';
import IconFontAwesome from 'react-native-vector-icons/FontAwesome';
import Input from '../../components/Input';
import Button from '../../components/Button';
import { getAvatars, getToken, authApi } from '../../services/api';
import { SvgUri } from 'react-native-svg';
import { decodeJWT } from '../../services/decodeJWT';
import { useNavigation } from '@react-navigation/native';
import { removeToken } from '../../services/api';

// This screen is responsible for displaying the user's profile information
export default function ProfileInformationScreen({ navigation }) {
  const [modalVisible, setModalVisible] = useState(true);
  const [avatarModalVisible, setAvatarModalVisible] = useState(false);
  const [avatars, setAvatars] = useState([]);
  const [loading, setLoading] = useState(false);
  const [userId, setUserId] = useState(null);
  const [userData, setUserData] = useState({});
  const [selectedAvatarUrl, setSelectedAvatarUrl] = useState('');
  const navigationUse = useNavigation();

  // Function to close the modal and navigate to the Home screen
  const handleClose = () => {
    setModalVisible(false);
    navigation.navigate('Home');
  };

  // Function to open the avatar selection modal
  const openAvatarModal = () => {
    setAvatarModalVisible(true);
  };

  // Fetch avatars when the screen is loaded
  useEffect(() => {
    const loadData = async () => {
      const token = await getToken();
      if (token) {
        const decodedToken = decodeJWT(token);
        setUserId(decodedToken.id);
        console.log('User ID:', decodedToken.id);

        try {
          const response = await authApi.get(`/users/${decodedToken.id}`);
          setUserData(response.data);
          console.log('User data:', response.data);
        } catch (error) {
          console.error('Failed to fetch user data:', error);
        }
      }
      fetchAvatars();
    };

    loadData();

    const timer = setTimeout(() => {
      console.log('ProfileInformationScreen loaded');
    }, 2000);

    return () => clearTimeout(timer);
  }, []);

  // Function to fetch avatars from the API
  const fetchAvatars = async () => {
    setLoading(true);
    try {
      const data = await getAvatars();
      setAvatars(
        Object.entries(data).map(([key, value]) => ({
          id: key,
          url: value
        }))
      );
      console.log('Avatares carregados:', avatars);
    } catch (error) {
      console.error('Erro ao buscar avatares:', error);
    } finally {
      setLoading(false);
    }
  };

  // Function to save the user profile
  const handleSave = async () => {
    try {
      const updatedData = {
        ...userData,
        image: selectedAvatarUrl
      };

      const response = await authApi.put(`/users/${userId}`, updatedData);
      console.log('Perfil atualizado com sucesso:', response.data);
      Alert.alert('Sucesso', 'Perfil atualizado com sucesso.');
      navigation.navigate('Profile');
    } catch (error) {
      console.error('Erro ao atualizar perfil:', error);
      Alert.alert('Erro', 'Falha ao atualizar o perfil. Tente novamente.');
    }
  };

  // Function redirecting to the Lgin screen
  const handleLogout = async () => {
    await removeToken(); 
    navigationUse.reset({
      index: 0,
      routes: [{ name: 'Login' }]
    });
  };

  return (
    <View>
      <ScrollView
        style={styles.modalView}
        contentContainerStyle={{ alignItems: 'center' }}
      >
        <TouchableOpacity style={styles.closeButton} onPress={handleClose}>
          <Text style={styles.closeButtonText}>
            <IconClose name='close' size={30} color='#000' />
          </Text>
        </TouchableOpacity>
        {userData.image ? (
          <SvgUri width='120' height='120' uri={userData.image} />
        ) : (
          <IconFontAwesome name='user-circle-o' size={100} color='#000' />
        )}
        <Button
          label='Selecionar Avatar'
          style={styles.buttonIcon}
          onPress={openAvatarModal}
        />
        <View style={styles.inputContainer}>
          <Input
            label={'Nome Completo'}
            value={userData.name}
            style={styles.input}
          />

          <Input
            label={'Nome da Empresa'}
            value={userData.company_name}
            style={styles.input}
          />

          <Input label={'Cargo'} value={userData.office} style={styles.input} />

          <Input label={'Email'} value={userData.email} style={styles.input} />

          <Input
            label={'Perfil do Linkedin'}
            placeholder={'Perfil do Linkedin'}
            value={userData.linkedin_link}
          />
        </View>

        <Modal
          animationType='slide'
          transparent={true}
          visible={avatarModalVisible}
          onRequestClose={() => setAvatarModalVisible(false)}
        >
          <View style={styles.modalContainer}>
            <TouchableOpacity style={styles.closeButton} onPress={handleClose}>
              <Text style={styles.closeButtonText}>
                <IconClose name='close' size={30} color='#FFFF' />
              </Text>
            </TouchableOpacity>
            {loading ? (
              <View style={styles.loadingContainer}>
                <ActivityIndicator size='large' color='#0000ff' />
              </View>
            ) : (
              <ScrollView style={styles.scrollView}>
                <View style={styles.avatarContainer}>
                  {avatars.map(avatar => (
                    <TouchableOpacity
                      key={avatar.id}
                      style={styles.avatar}
                      onPress={() => {
                        console.log('Avatar selecionado:', avatar.url);
                        setSelectedAvatarUrl(avatar.url);
                        setAvatarModalVisible(false);
                      }}
                    >
                      <SvgUri width='100' height='100' uri={avatar.url} />
                    </TouchableOpacity>
                  ))}
                </View>
              </ScrollView>
            )}
          </View>
        </Modal>
        <Button label='Salvar' style={styles.button} onPress={handleSave} />
        <Button
          label='Sair do app'
          style={styles.buttonCloseApp}
          onPress={handleLogout}
        />
      </ScrollView>
    </View>
  );
}
