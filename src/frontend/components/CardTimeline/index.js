import React from 'react';
import { View, Text, TouchableOpacity } from 'react-native';
import styles from './styles';
import IconGreen from '../IconGreen';
import TextArea from '../TextArea';

// This component is a card used on the timeline screen to show information about projects that have been connected
export default function CardTimeline({ title, nameUser, companyUser, imageUser, value, nameConnection }) {
    return(
        <View style={styles.container}>
            <View>
                <Text style={styles.container__title}>{title}</Text>

            </View>
            <View style={styles.container__iconGreen}>
                <IconGreen
                    image={imageUser}
                    name={nameUser}
                    company={companyUser}
                />
                <IconGreen
                    image={imageUser}
                    name={nameConnection}
                    company={companyUser}
                    background={false}
                />
            </View>
            <TextArea
            title={"Observações do Usuário"}
            minHeight={80}
            value={value}
            />

        </View>
    )
}
