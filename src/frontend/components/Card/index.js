import React from 'react';
import { View, Text, TouchableOpacity } from 'react-native';
import styles from './styles';
import IconGreen from '../IconGreen';
import IconApprove from 'react-native-vector-icons/AntDesign';
import IconReject from 'react-native-vector-icons/AntDesign';
import IconEdit from 'react-native-vector-icons/Feather';
import IconTrash from 'react-native-vector-icons/Feather';

// This component is a Card that can be used to display project information, notifications and user projects
export default function Card({title, description, nameUser, companyUser, imageUser, notification, MyProject, navigation, projectId, handleDeleteProject, handleEditProject }) {

  // Function that will be called when an icon is pressed
  const handlePress = (iconName) => {
    if (iconName === 'card') {
      console.log('apertou', projectId)
      navigation.navigate('Project', { id: projectId }); 
    } else {
      console.log(`${iconName} pressed`);
    }
  };

  // This component is the content of the card
  const CardContent = () => (
    <View style={styles.container}>
      <View style={(notification || MyProject) ? styles.informationsNotification : styles.informations}>
        <View>
          <Text style={styles.container__title}>{title}</Text>
          <Text style={styles.container__description}>{description}</Text>
        </View>
        {(notification || MyProject) && (
          <View style={styles.container__actions}>
            {notification && (
              <>
                <TouchableOpacity onPress={() => handlePress('Icon approve')}>
                  <IconApprove style={styles.IconApprove} name="check" size={30} color="#80DB55" />
                </TouchableOpacity>
                <TouchableOpacity onPress={() => handlePress('Icon reject')}>
                  <IconReject name="close" size={30} color="#000" />
                </TouchableOpacity>
              </>
            )}
            {MyProject && (
              <View style={styles.containerMyProject}>
                <TouchableOpacity onPress={handleDeleteProject}>
                  <IconTrash style={styles.IconApprove} name="trash-2" size={30} color="#BB3F56" />
                </TouchableOpacity>
                <TouchableOpacity onPress={handleEditProject}>
                  <IconEdit name="edit-2" size={25} color="#000" />
                </TouchableOpacity>
              </View>
            )}
          </View>
        )}
      </View>
      {!MyProject && (
        <View style={(notification || MyProject) ? styles.container__iconGreenNotification : styles.container__iconGreen}>
          <IconGreen 
            image={imageUser}
            name={nameUser}
            company={companyUser}
          />
        </View>
      )}
    </View>
  );

  // If the card is a notification or a user project, it will not be clickable
  if (notification || MyProject) {
    return <View>
      <CardContent />
    </View>;
  } else {
    return <TouchableOpacity onPress={() => handlePress('card')} style={{ flex: 1 }}>
      <CardContent />
    </TouchableOpacity>;
  }
}
