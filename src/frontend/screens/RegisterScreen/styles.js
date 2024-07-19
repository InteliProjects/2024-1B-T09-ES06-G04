import { StyleSheet, Dimensions } from 'react-native';

const screenWidth = Dimensions.get('window').width; 
const screenHeight = Dimensions.get('window').height; 

const styles = StyleSheet.create({
  centeredView: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center'
  },
  fullScreenModalView: {
    width: screenWidth, 
    height: screenHeight, 
    backgroundColor: "white",
    padding: 12,
    paddingTop: 0,
    shadowColor: "#000",
    shadowOffset: {
      width: 0,
      height: 2
    },
    shadowOpacity: 0.25,
    shadowRadius: 4,
    elevation: 5
  },
  closeButton: {
    alignSelf: 'flex-end',
    marginLeft: 20,
    marginTop: 20,
    marginBottom: 20,
  },
  input: {
    width: '100%', 
    height: 40,
    marginBottom: 12,
    borderWidth: 1,
    padding: 10,
  },
  inputContainer: {
    width: 360,
    marginBottom: 30,
  },
  button: {
    borderRadius: 10,
    padding: 10,
    elevation: 2,
    backgroundColor: '#B6E99E',
    marginBottom: 40,
    width: '70%',
    shadowColor: "#000",
    shadowOffset: {
      width: 0,
      height: 2
    },
    shadowOpacity: 0.25,
    shadowRadius: 4,
    elevation: 5
  },
  textStyle: {
    fontSize: 20,
    color: "#464646",
    fontWeight: "light",
    textAlign: "center",
  },
  closeButtonText: {
    fontSize: 40,
    fontWeight: 'bold',
  },
  profile: {
    marginBottom: 20,
  }
});

export default styles;