import { StyleSheet } from 'react-native';

const styles = StyleSheet.create({
  closeButton: {
    alignSelf: 'flex-end',
    marginLeft: 20,
    marginTop: 20,
    marginBottom: 20,
    marginRight: 20
  },
  button: {
    backgroundColor: '#B6E99E',
    borderRadius: 10,
    padding: 10,
    margin: 10,
    width: 200,
    height: 50,
    alignItems: 'center'
  },
  buttonIcon: {
    width: 200,
    fontSize: 15,
    marginBottom: 32
  },
  inputContainer: {
    alignItems: 'center'
  },
  input: {
    marginBottom: 32
  },
  modalContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: 'rgba(0, 0, 0, 0.5)'
  },
  loadingContainer: {
    justifyContent: 'center',
    alignItems: 'center',
    height: '100%'
  },
  avatarContainer: {
    flexDirection: 'row',
    flexWrap: 'wrap',
    justifyContent: 'center',
    padding: 20
  },
  avatar: {
    margin: 5
  },
  buttonCloseApp: {
    backgroundColor: '#ff6969',
    color: '#fff',
    marginBottom: 10,
  }
});

export default styles;
