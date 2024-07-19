import { StyleSheet } from 'react-native';

export default StyleSheet.create({
  container: {
    display: 'flex',
    flexDirection: 'row',
    alignItems: 'center',
    backgroundColor: '#A3BFB7',
    maxWidth: '40%',
    borderRadius: 10,
    padding: 7,
    gap: 10,
  },
  company: {
    marginTop: 5,
    fontSize: 9,
    fontWeight: '300',
    maxWidth: '90%',
  },
  name: {
    maxWidth: '90%',
  }
});