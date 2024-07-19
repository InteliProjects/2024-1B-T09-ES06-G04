import { Button, StyleSheet } from 'react-native';

const styles = StyleSheet.create({
  outerContainer: {
    flex: 1,
    backgroundColor: '#fff',
  },
  container: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
    zIndex: 1, 
  },
  logo: {
    marginBottom: 10,

  },

  titleText: {
    fontSize: 24,
    fontWeight: '300',
    marginBottom: 60,
    color: '#BB3F56',
  },
  inputContainer: {
    width: '80%',
  },
  inputWrapper: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'center',
    marginBottom: 16,
  },
  input: {
    width: 330,
    height: 40,
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

  blueCircle: {
    position: 'absolute',
    top: '42%',
    left: '-5%',
    width: 700,
    height: 800,
    backgroundColor: '#ABC4DC',
    borderRadius: 700,
    zIndex: 0, 
  },

  loginButton: {
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

  registerButton: {
    position: 'absolute',
    bottom: '5%',
    alignItems: 'center',
    justifyContent: 'center',
    width: 140,
    height: 40,
  },

  buttonText: {
    fontSize: 20,

  },

  signupText: {
    position: 'absolute',
    bottom: '5%',

  }
});


export default styles;
