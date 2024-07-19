import React, { useEffect, useState } from 'react';
import { TouchableOpacity, View, Image, Alert, TextInput } from 'react-native';
import IconFontAwesome from "react-native-vector-icons/FontAwesome";
import IconAntDesign from "react-native-vector-icons/AntDesign";
import AsyncStorage from '@react-native-async-storage/async-storage';
import { authApi } from '../../services/api';
import styles from './styles';
import { SvgUri } from 'react-native-svg';
import { useNavigation } from '@react-navigation/native'; 
import IconSearch from "react-native-vector-icons/AntDesign";
import { decodeJWT } from '../../services/decodeJWT';

// This is the Header component, which is responsible for displaying the header of the application
export default function Header() {
  const [avatarUrl, setAvatarUrl] = useState('');
  const navigation = useNavigation();
  const [searchText, setSearchText] = useState(''); 

  // Function to handle the press of the icons
  const handlePress = iconName => {
    console.log(`${iconName} pressed`);

    if (iconName === 'my projects') {
      navigation.navigate('MyProjects');
    } else if (iconName === 'profile') {
      navigation.navigate('Profile');
    } else if (iconName === 'search') {
      navigation.navigate('SearchResults', { searchText });
    }
  };

  // Fetch the user profile when the screen is loaded
  useEffect(() => {
    const loadUserProfile = async () => {
      try {
        const token = await AsyncStorage.getItem('authToken');
        if (!token) {
          console.log('No token found');
          return;
        }

        const decodedToken = decodeJWT(token);

        const response = await authApi.get(`/users/${decodedToken.id}`);
        if (response.data) {
          setAvatarUrl(response.data.image);
          console.log(avatarUrl)
        }
      } catch (error) {
        console.error('Failed to fetch user profile:', error);
        Alert.alert(
          'Erro',
          'Não foi possível buscar as informações do usuário.'
        );
      }
    };

    loadUserProfile();
  }, []);

  return (
    <View style={styles.container}>
      <TouchableOpacity onPress={() => handlePress('profile')}>
        {avatarUrl ? (
          <SvgUri width='40' height='40' uri={avatarUrl} style={styles.avatarIcon}/>
        ) : (
          <IconFontAwesome name='user-circle-o' size={40} color='#000' />
        )}
      </TouchableOpacity>

      <View style={styles.searchContainer}>
        <TextInput
          style={styles.searchInput}
          placeholder="Pesquise projetos ou CEO's"
          onChangeText={setSearchText} 
          value={searchText} 
        />
        <TouchableOpacity onPress={() => handlePress('search')}>
          <IconSearch name='search1' size={20} color='#a6a6a6' />
        </TouchableOpacity>
      </View>

      <TouchableOpacity onPress={() => handlePress('my projects')}>
        <IconAntDesign name='folderopen' size={40} color='#E3B146' />
      </TouchableOpacity>
    </View>
  );
}
