import { StyleSheet } from 'react-native';

const styles = StyleSheet.create({
  title: {
    fontSize: 20,
    fontWeight: '300',
    textAlign: 'left',
    marginBottom: 5,
    marginTop: 10,
    width: '100%',
    alignSelf: 'left',
  },
  input: {
    fontSize: 16,
    padding: 12,
    backgroundColor: '#FAFAFA',
    textAlignVertical: 'top',
    borderRadius: 10,
    shadowColor: '#000',
    shadowOffset: {
      width: 0,
      height: 2,
    },
    shadowOpacity: 0.25,
    shadowRadius: 3.84,
    elevation: 5,
  },
});

export default styles;
