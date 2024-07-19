import { TouchableOpacity, Text } from 'react-native';
import styles from './styles';

// Define a Button component that accepts label, style, and other props
export default function Button({ label, style, ...props }) {
  return (
    <TouchableOpacity style={[styles.button, style]} {...props}>
      <Text style={styles.buttonText}>{label}</Text>
    </TouchableOpacity>
  );
}
