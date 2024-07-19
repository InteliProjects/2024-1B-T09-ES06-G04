import { StyleSheet } from 'react-native';

export default StyleSheet.create({
  container: {
    paddingEnd: 16,
    paddingStart: 16,
    backgroundColor: '#fff',
  },
  modalView: {
    padding: 20,
    marginTop: 40,
    alignItems: 'center',
  },
  closeButton: {
    position: 'absolute',
    top: -15,
    right: 10,
  },
  saveButton: {
    backgroundColor: '#B6E99E',
    borderRadius: 10,
    width: 200,
    height: 50,
    alignItems: 'center', 
    justifyContent: 'center' 
  },
  descriptionInput: {
    width: 20,
  }
});
