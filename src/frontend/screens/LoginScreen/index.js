import React, { useState } from 'react';
import {
  ScrollView,
  View,
  TouchableOpacity,
  Text,
  Image,
  Alert
} from 'react-native';
import styles from './styles';
import { useNavigation } from '@react-navigation/native';
import { login } from '../../services/api';
import Input from '../../components/Input';
import Button from '../../components/Button';
import {decodeJWT} from '../../services/decodeJWT';

// The LoginScreen component is responsible for rendering the login screen
export default function LoginScreen() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const navigation = useNavigation();

  // The handleLogin function is responsible for handling the login process
  const handleLogin = async () => {
    try {
      const token = await login(email, password);
      console.log('Token:', token);

      const decodedToken = decodeJWT(token);
      console.log('User ID:', decodedToken.id);

      navigation.navigate('Home');
    } catch (error) {
      console.error('Login failed:', error);
      Alert.alert(
        'Erro de Login',
        'Falha no login, verifique suas credenciais e tente novamente.'
      );
    }
  };

  return (
    <View style={styles.outerContainer}>
      <View style={styles.blueCircle} />
      <ScrollView contentContainerStyle={styles.container}>
        <View style={styles.logoContainer}>
          <Image
            source={require('../../assets/logoFDC.png')}
            style={styles.logo}
          />
        </View>
        <Text style={styles.titleText}>CLI | CEO'S LEGACY</Text>
        <View style={styles.inputContainer}>
          <View style={styles.inputWrapper}>
            <Input placeholder='Email' value={email} onChangeText={setEmail} />
          </View>
          <View style={styles.inputWrapper}>
            <Input
              placeholder='Senha'
              value={password}
              onChangeText={setPassword}
              secureTextEntry={true}
              style={{ marginTop: 10 }}
            />
          </View>
        </View>
        <Button label='Entrar' onPress={handleLogin} />

        <TouchableOpacity
          style={styles.registerButton}
          onPress={() => navigation.navigate('Register')}
        >
          <Text style={styles.signupText}>É novo? Cadastre-se</Text>
        </TouchableOpacity>
      </ScrollView>
    </View>
  );
}
