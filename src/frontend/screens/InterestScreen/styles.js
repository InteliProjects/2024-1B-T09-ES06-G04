import { StyleSheet } from 'react-native';

const styles = StyleSheet.create({
  containerTitle: {
    marginTop: 20,
    marginBottom: 20,
  },
  container: {
    flex: 1,
    padding: 10,
    alignItems: 'center',
    width: '100%',  
  },
  item: {
    padding: 10,
    marginBottom: 30,
    backgroundColor: '#f0f0f0',
    borderRadius: 20,
    borderWidth: 0.5,
    width: '100%',
    marginLeft: 30,
    marginRight: 30
  },
  selectedItem: {
    padding: 10,
    marginBottom: 30,
    backgroundColor: '#E2E0E0',
    borderRadius: 20,
    borderWidth: 12,
    width: '100%',
    marginLeft: 30,
    marginRight: 30
  },
  text: {
    fontSize: 16
  },
  selectedText: {
    fontSize: 16,
    fontWeight: '500'
  },
});

export default styles;