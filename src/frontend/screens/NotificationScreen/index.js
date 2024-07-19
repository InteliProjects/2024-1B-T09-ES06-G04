import React from 'react';
import { ScrollView, Text } from 'react-native';
import Card from '../../components/Card';
import styles from './styles';
import Title from '../../components/Title';

// This screen is responsible for displaying the user's notifications
export default function NotificationScreen() {
  return (
    <ScrollView style={styles.container}>
      <Title 
        title="Convites"
      />
      <Text> Nenhuma notificação </Text>
    </ScrollView>
  );
}
