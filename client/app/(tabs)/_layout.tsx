import { StyleSheet,Image, Text, View } from "react-native";
import React from 'react'
import {Tabs, Redirect } from 'expo-router'
import { icons } from "../../constants";

interface TabIconProps {
  icon: any;
  color: string;
  name: string;
  focused: boolean;
}


const TabIcon: React.FC<TabIconProps> = ({ icon, color, name, focused }) => {
  return (
    <View style={{ flexDirection: 'row', alignItems: 'center', justifyContent: 'center', gap: 2 }}>
      <Image
        source={icon}
        resizeMode="contain"
        style={{ width: 24, height: 24, tintColor: color }}
      />
    </View>
  );
};




const TabsLayout = () => {

  return (
    <>
        <Tabs
          screenOptions={{
            tabBarShowLabel:false,
            tabBarActiveTintColor:'black'
          }}
        >
          <Tabs.Screen
            name="Home"
            options={{
              title: "",
              headerShown: false,
              tabBarIcon: ({ color, focused }) => (
                <TabIcon
                  icon={icons.home}
                  color={color}
                  name="Home"
                  focused={focused}
                />
              ),
            }}
          />


          <Tabs.Screen
            name="SearchScreen"
            options={{
              title: "",
              headerShown: false,
              tabBarIcon: ({ color, focused }) => (
                <TabIcon
                  icon={icons.search}
                  color={color}
                  name="Search"
                  focused={focused}
                />
              ),
            }}
            listeners={({ navigation, route }) => ({
              tabPress: (e) => {
                const state = navigation.getState();
                const currentRoute = state.routes[state.index].name;
          
                if (currentRoute === "SearchScreen") {
                  // Clear params only if we're already on SearchScreen
                  navigation.navigate("SearchScreen");
                } else {
                  // Prevent default behavior and navigate to SearchScreen
                  e.preventDefault();
                  const currentParams = route.params || {};
                  navigation.navigate("SearchScreen", { ...currentParams });
        
                }
              },
            })}
          />

          <Tabs.Screen
            name="PostScreen"
            options={{
              title: "",
              headerShown: false,
              tabBarIcon: ({ color, focused }) => (
                <TabIcon
                  icon={icons.plus}
                  color={color}
                  name="Post"
                  focused={focused}
                />
              ),
            }}
          />

          
          <Tabs.Screen
            name="ClipScreen"
            options={{
              title: "",
              headerShown: false,
              tabBarIcon: ({ color, focused }) => (
                <TabIcon
                  icon={icons.clip}
                  color={color}
                  name="Clip"
                  focused={focused}
                />
              ),
            }}
          />

          <Tabs.Screen
            name="ProfileScreen"
            options={{
              title: "",
              headerShown: false,
              tabBarIcon: ({ color, focused }) => (
                <TabIcon
                  icon={icons.profile}
                  color={color}
                  name="Profile"
                  focused={focused}
                />
              ),
            }}
          />


        </Tabs>
    </>
  )
}

export default TabsLayout

const styles = StyleSheet.create({})