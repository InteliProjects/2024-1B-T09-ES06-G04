import { ScrollView, Text, TouchableOpacity, View, Alert } from 'react-native';
import Title from '../../components/Title';
import React, { useState } from 'react';
import styles from './styles';
import Button from '../../components/Button';
import { updateUserInterests } from '../../services/api';


const borderColors = ['#4E5F3D', '#7D8873', '#81859E', '#454861', '#DF9D4B', '#ECC087'];

// Function to render the RegisterScreen component
export default function RegisterScreen({navigation, route}) {

  const { userId } = route.params;
  const [selectedInterests, setSelectedInterests] = useState([]);

  // Function to handle the form submission
  const handleSubmit = async () => {
    if (selectedInterests.length < 3) {
      Alert.alert('Atenção', 'Escolha no mínimo três interesses para prosseguir.');
    } else {
      try {
        await updateUserInterests(userId, selectedInterests);  
        Alert.alert('Sucesso', 'Interesses cadastrados com sucesso! Realize o login com suas credenciais para acessar a plataforma.');
        navigation.navigate('Login');
      } catch (error) {
        Alert.alert('Erro', 'Falha ao atualizar interesses');
      }
    }
  }

  const interests = [
    "Redução do Impacto Ambiental",
    "Conservação do planeta",
    "Integridade e Práticas Éticas",
    "Produtividade e Competitividade",
    "Diversidade & Inclusão",
    "Bem Estar, Saúde e Felicidade"
  ];


  // Function to handle the selection of interests
  const handleSelectInterest = (interest) => {
    if (selectedInterests.includes(interest)) {
      setSelectedInterests(selectedInterests.filter(item => item !== interest));
    } else {
      setSelectedInterests([...selectedInterests, interest]);
    }
  };

  // Function to get the border color for the interest item
  const getBorderColor = (index) => borderColors[index % borderColors.length];

  return (
    <>
      <View style={styles.containerTitle}>
        <Title 
          title="Quais são os seus interesses?"
          subtitle="Escolha no mínimo três para prosseguir"
        />
      </View>
      <View style={styles.container}>
        <ScrollView style={{width: '100%'}} contentContainerStyle={{ alignItems: 'center'}}>
          {interests.map((interest, index) => (
            <TouchableOpacity
              key={index}
              style={[
                selectedInterests.includes(interest) ? styles.selectedItem : styles.item,
                { borderColor: getBorderColor(index), borderWidth: 0.6 }
              ]}
              onPress={() => handleSelectInterest(interest)}
            >
              <Text style={selectedInterests.includes(interest) ? styles.selectedText : styles.text}>{interest}</Text>
            </TouchableOpacity>
          ))}
        </ScrollView>

        <Button 
          label="Prosseguir" 
          onPress={handleSubmit} 
          style={{position: 'absolute', bottom: 100}}
        />

      </View>
    </>
  );
};
