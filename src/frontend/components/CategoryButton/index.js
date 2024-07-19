import { TouchableOpacity, Text } from 'react-native';
import styles from './styles';

// This component is a button that can be used to filter projects by category
export default function CategoryButton({ title, onPress, backgroundColor, color }) {
  return (
    <TouchableOpacity style={[styles.button, { backgroundColor: backgroundColor}]} onPress={onPress}>
      <Text style={[styles.buttonText, {color: color} ]}>{title}</Text>
    </TouchableOpacity>
  );
}