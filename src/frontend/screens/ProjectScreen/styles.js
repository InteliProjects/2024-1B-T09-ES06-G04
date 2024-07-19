import { StyleSheet } from "react-native";

export default StyleSheet.create({
  outerContainer: {
    flex: 1,
    backgroundColor: '#f5f5f5',
  },
  container: {
    flexGrow: 1,

  },
  header: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: 20,
    paddingRight: 12,
  },

  coverPhotoContainer: {
    alignItems: 'center',
    marginBottom: 20,
    padding: 12,
  },
  coverPhoto: {
    width: '100%',
    height: 200,
    resizeMode: 'cover',
    borderRadius: 10,
    backgroundColor: '#e0e0e0',
    padding: 12,
  },
  connectButtonContainer: {
    flexDirection: 'row',
    alignItems: 'flex-start', // Alterado para alinhar à esquerda
    justifyContent: 'space-between',
    marginBottom: 16,
    padding: 12,
  },
  connectButton: {
    flexDirection: 'row',
    backgroundColor: '#fff',
    padding: 12,
    borderRadius: 5,
    alignItems: 'center',
    borderWidth: 1,
    borderColor: '#e0e0e0',
    borderRadius: 10,
  },
  connectButtonIcon: {
    marginRight: 5,
  },
  connectButtonText: {
    fontSize: 16,
    color: '#000',
  },
  descriptionContainer: {
    marginBottom: 16,
    paddingLeft: 12,
    paddingRight: 12,
  },
  sectionTitle: {
    fontSize: 16,
    fontWeight: 'bold',
    marginBottom: 12,
  },
  description: {
    fontSize: 14,
    color: '#333',
    lineHeight: 22,
  },
  connectionsContainer: {
    marginBottom: 16,
    padding: 12,
  },
  connection: {
    flexDirection: 'row',
    alignItems: 'center',
    marginBottom: 12,
    padding: 10,
    backgroundColor: '#E4EEEB',
    borderRadius: 10,
  },
  
  connectionImage: {
    marginRight: 10
  }

});
