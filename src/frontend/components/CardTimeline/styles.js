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
        display: 'flex',
        flexDirection: 'column'
    },
        container__title: {
            fontSize: 24,
            fontWeight: '300',
            marginBottom: 16,
        },

        container__iconGreen: {
            position: 'relative',
            flexDirection: 'row',
            gap: 12, 
        },
});
