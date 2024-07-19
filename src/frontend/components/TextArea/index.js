import React from 'react';
import { View, TextInput, Text } from 'react-native';
import styles from './styles';

// This component is responsible for presenting a text area that can be used to insert text
export default function TextArea({ title, minHeight, value, onChangeText }) {
  return (
    <View style={styles.container}>
      <Text style={styles.title}>{title}</Text>
      <TextInput
        value={value}
        style={[styles.input, { minHeight: minHeight, minWidth: '95%', maxWidth: '95%'}]}
        onChangeText={onChangeText}
        multiline={true}
      />
    </View>
  );
}