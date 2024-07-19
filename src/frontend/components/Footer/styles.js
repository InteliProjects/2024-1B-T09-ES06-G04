import { StyleSheet } from 'react-native';

export default StyleSheet.create({
  container: {
    position: 'fixed',
    bottom: 0,
    backgroundColor: '#ffff',
    padding: 10,
    paddingLeft: 10,
    paddingRight: 10,
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    width: '100%',
    shadowColor: '#000',
    shadowOffset: {
      width: 0,
      height: 2,
    },
    shadowOpacity: 0.25,
    shadowRadius: 3.84,
    elevation: 5,
    zIndex: 1,
  },
  text: {
    fontSize: 10,
    marginTop: 2,
    textAlign: 'center',
  }, 
  button: {
    display: 'flex',
    alignItems: 'center',
    marginLeft: 0,
    marginRight: 0,
    width: 50,
  },
});
