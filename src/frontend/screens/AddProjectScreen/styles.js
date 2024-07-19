import { StyleSheet } from 'react-native';

export default StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#fff',
      },
      scrollContent: {
        flexGrow: 1,
        padding: 12,
        paddingTop: 0,
      },
    header: {
      flexDirection: 'row',
      justifyContent: 'space-between',
      alignItems: 'center',
    },
    buttonContainer: {
      marginBottom: 12,
    },
    closeButton:{
      marginTop: 10,
      
    }
    });