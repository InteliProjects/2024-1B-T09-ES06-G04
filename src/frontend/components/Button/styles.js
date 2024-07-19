import { StyleSheet } from 'react-native';

const styles = StyleSheet.create({
  button: {
    marginTop: 24,
    alignItems: 'center',
    justifyContent: 'center',
    width: 140,
    height: 40,
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
  buttonText: {
    fontSize: 20,
  }
});

export default styles;