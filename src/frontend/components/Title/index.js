import { Text, View } from 'react-native';
import styles from './styles';


// Function that when called defines a title and subtitle based on the passed parameter
export default function Title({ title, subtitle }) {
    return (
        <View style={styles.container}>
            <Text style={styles.title}>
                {title}
            </Text>
            {subtitle ? (
                <Text style={styles.subtitle}>
                    {subtitle}
                </Text>
            ) : null}
        </View>
    );
};