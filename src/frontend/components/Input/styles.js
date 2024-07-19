import { StyleSheet } from 'react-native';

const styles = StyleSheet.create({
  input: {
    width: 330,
    height: 50,
    padding: 10,
    borderRadius: 10,
    backgroundColor: '#FAFAFA', 
    shadowColor: '#000',
    shadowOffset: {
      width: 0,
      height: 2,
    },
    shadowOpacity: 0.50,
    shadowRadius: 1,
    elevation: 5,
  },
  label: {
    fontSize: 18,
    marginBottom: 5,
    fontWeight: 'lighter',
    letterSpacing: 0.5,
    color: "#444444"
  }
});

export default styles;