import { View, Text, Image } from "react-native";
import styles from "./styles";

// This component is responsible for presenting basic user informationexport default function IconGreen({ image, name, company, background }) {
export default function IconGreen({ image, name, company, background }) {

  // The containerStyle variable is responsible for defining the style of the container that will contain the user's information.
  const containerStyle = [
    styles.container,
    { backgroundColor: background === false ? 'transparent' : '#A3BFB7' }
  ];

  return (
    <View style={containerStyle}>
      <Image style={styles.image} source={image} />
      <View style={styles.informations}>
        <Text style={styles.name}>{name}</Text>
        <Text style={styles.company}>{company}</Text>
      </View>
    </View>
  );
}
