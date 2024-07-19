import { View, Text, TextInput } from 'react-native';
import styles from './styles';

// The Input component is responsible for rendering an input field
export default function Input({ label, style, ...props }) {
  return (
    <View style={styles.inputContainer}>

      {label && <Text style={styles.label}>{label}</Text>}

      <TextInput 
        style={[styles.input, style]} 
        {...props} 
      />
    </View>
  );
}