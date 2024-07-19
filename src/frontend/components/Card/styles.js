import { StyleSheet } from 'react-native';

export default StyleSheet.create({
  container: {
    marginLeft: 12,
    marginRight: 12,
    marginTop: 12,
    marginBottom: 12,
    backgroundColor: '#FAFAFA',
    padding: 24,
    borderRadius: 10,
    shadowColor: '#000',
    shadowOffset: {
      width: 0,
      height: 2,
    },
    shadowOpacity: 0.50,
    shadowRadius: 1,
    elevation: 5,
  },
  informations: {
    marginBottom: 10,
  },
  container__title: {
    fontSize: 24,
    fontWeight: '300',
    marginBottom: 1,
    maxWidth: 220
  },
  container__description: {
    fontSize: 16,
    color: '#666',
    maxWidth: 230
  },
  container__iconGreen: {
    position: 'relative',
    left: 210,
    bottom: -10,
  },
  container__iconGreenNotification: {
    position: 'relative',
    left: 210,
    bottom: -10,
    flexDirection: 'row',
    alignItems: 'center',
    left: 0,
    marginTop: 20,
  },
  container__actions: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
    marginLeft: 10,
  },
  informationsNotification : {
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'space-between',
    width: '100%'
  }, 
  IconApprove: {
    marginRight: 15,
  },
  containerMyProject: {
    display: 'flex',
    flexDirection: 'row',
  }
});
