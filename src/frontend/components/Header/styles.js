import { StyleSheet } from 'react-native';

export default StyleSheet.create({
  container: {
    position: 'fixed',
    top: 0,
    backgroundColor: '#ffff',
    padding: 10,
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
  searchContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    backgroundColor: '#f8f8f8',
    border: 'none',
    borderRadius: 20,
    flex: 1,
    paddingHorizontal: 10, 
    marginHorizontal: 20,
    marginLeft: 20,
    marginRight: 20,
    height: 30
  },
  searchInput: {
    flex: 1,
    paddingHorizontal: 10
  }
});