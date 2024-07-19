import React from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import HomeScreen from '../screens/HomeScreen';
import RecommendScreen from '../screens/RecommendScreen';
import TimelineScreen from '../screens/TimelineScreen';
import AddProjectScreen from '../screens/AddProjectScreen'; 
import MyProjectsScreen from '../screens/MyProjectsScreen';
import NotificationScreen from '../screens/NotificationScreen';
import ProjectScreen from '../screens/ProjectScreen';
import LoginScreen from '../screens/LoginScreen';
import Layout from '../layout';
import RegisterScreen from '../screens/RegisterScreen';
import InterestScreen from '../screens/InterestScreen';
import SearchResultsScreen from '../screens/SearchResultsScreen';
import ProfileScreen from '../screens/ProfileInformationScreen';

// This function is responsible for defining the routes of the application
const Stack = createNativeStackNavigator();

// The Routes component is responsible for defining the routes of the application
export default function Routes() {
  return (
    <NavigationContainer>
      <Stack.Navigator
        initialRouteName="Login"
        screenOptions={{
          headerShown: false, 
        }}
      >
        
        <Stack.Screen name="Home">
          {props => (
            <Layout navigation={props.navigation}>
              <HomeScreen {...props} />
            </Layout>
          )}
        </Stack.Screen>
        <Stack.Screen name="Recommend">
          {props => (
            <Layout navigation={props.navigation}>
              <RecommendScreen {...props} />
            </Layout>
          )}
        </Stack.Screen>
        <Stack.Screen name="Timeline">
          {props => (
            <Layout navigation={props.navigation}>
              <TimelineScreen {...props} />
            </Layout>
          )}
        </Stack.Screen>
        <Stack.Screen name="AddProject">
          {props => (
            <Layout navigation={props.navigation} hideHeaderFooter={true}>
              <AddProjectScreen {...props} />
            </Layout>
          )}
        </Stack.Screen>
        <Stack.Screen name="MyProjects">
          {props => (
            <Layout navigation={props.navigation}>
              <MyProjectsScreen {...props} />
            </Layout>
          )}
        </Stack.Screen>
        <Stack.Screen name="Notification">
          {props => (
            <Layout navigation={props.navigation}>
              <NotificationScreen {...props} />
            </Layout>
          )}
        </Stack.Screen>
        <Stack.Screen name="Project">
          {props => (
            <Layout navigation={props.navigation}>
              <ProjectScreen {...props} />
            </Layout>
          )}
        </Stack.Screen>
        <Stack.Screen name="Login">
          {props => (
            <Layout navigation={props.navigation} hideHeaderFooter={true}>
              <LoginScreen {...props} />
            </Layout>
          )}
        </Stack.Screen>
        <Stack.Screen name="Register">
          {props => (
            <Layout navigation={props.navigation} hideHeaderFooter={true}>
              <RegisterScreen {...props} />
            </Layout>
          )}
        </Stack.Screen>
        <Stack.Screen name="Interest">
          {props => (
            <Layout navigation={props.navigation} hideHeaderFooter={true}>
              <InterestScreen {...props} />
            </Layout>
          )}
        </Stack.Screen>
        <Stack.Screen name="SearchResults">
          {props => (
            <Layout navigation={props.navigation}>
              <SearchResultsScreen {...props} />
            </Layout>
          )}
        </Stack.Screen>
        <Stack.Screen name="Profile">
          {props => (
            <Layout navigation={props.navigation} hideHeaderFooter={true}>
              <ProfileScreen {...props} />
            </Layout>
          )}
        </Stack.Screen>
      </Stack.Navigator>
    </NavigationContainer>
  );
};