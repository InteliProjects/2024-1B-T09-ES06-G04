import { SafeAreaView, View } from 'react-native';
import Header from '../components/Header';
import Footer from '../components/Footer';
import styles from './styles';

// This component is responsible for defining the layout of the application
export default function Layout({ children, navigation, hideHeaderFooter }) {
  return (
    <SafeAreaView style={styles.container}>
      {!hideHeaderFooter && <Header />}
      <View style={styles.content}>
        {children}
      </View>
      {!hideHeaderFooter && <Footer navigation={navigation} />}
    </SafeAreaView>
  );
};
