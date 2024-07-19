import { View, TouchableOpacity, Text } from 'react-native';
import styles from './styles';
import IconHome from 'react-native-vector-icons/Ionicons';
import IconTimeline from 'react-native-vector-icons/MaterialCommunityIcons';
import IconAdd from 'react-native-vector-icons/AntDesign';
import IconRecommend from 'react-native-vector-icons/AntDesign';
import IconNotification from 'react-native-vector-icons/Ionicons';

// This component is the footer that appears on the bottom of the screen
export default function Footer({ navigation }) {

  // Function that will be called when an icon is pressed
  const handlePress = iconName => {
    console.log(`${iconName} pressed`);
    if (iconName === 'home') {
      navigation.navigate('Home');
    } else if (iconName === 'recommend') {
      navigation.navigate('Recommend');
    } else if (iconName === 'notification') {
      navigation.navigate('Notification');
    } else if (iconName == 'timeline') {
      navigation.navigate('Timeline');
    } else if (iconName == 'add') {
      navigation.navigate('AddProject');
    } else if (iconName == 'notification') {
      navigation.navigate('Notification');
    }
  };

  return (
    <View style={styles.container}>
      <TouchableOpacity
        style={styles.button}
        onPress={() => handlePress('home')}
      >
        <IconHome name='home-outline' size={30} color='#000' />
        <Text style={styles.text}>Início</Text>
      </TouchableOpacity>
      <TouchableOpacity
        style={styles.button}
        onPress={() => handlePress('timeline')}
      >
        <IconTimeline name='account-group-outline' size={30} color='#000' />
        <Text style={styles.text}>Timeline</Text>
      </TouchableOpacity>
      <TouchableOpacity
        style={styles.button}
        onPress={() => handlePress('add')}
      >
        <IconAdd name='plus' size={30} color='#BB3F56' />
        <Text style={styles.text}>Projeto</Text>
      </TouchableOpacity>
      <TouchableOpacity
        style={styles.button}
        onPress={() => handlePress('recommend')}
      >
        <IconRecommend name='like2' size={30} color='#000' />
        <Text style={styles.text}>Sugestões</Text>
      </TouchableOpacity>
      <TouchableOpacity
        style={styles.button}
        onPress={() => handlePress('notification')}
      >
        <IconNotification name='notifications-outline' size={30} color='#000' />
        <Text style={styles.text}>Convites</Text>
      </TouchableOpacity>
    </View>
  );
}
