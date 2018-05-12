//replace this with actual new post stuff
const Screen = (props) => (
   <View style={{ flex: 1, backgroundColor: '#fff', alignItems: 'center', justifyContent: 'center' }}>
     <Text>{props.title} Screen</Text>
   </View>
);

const NewPost = StackNavigator({
   NewPost: {
     screen: (props) => <Screen title="New Post" {...props} />,
     navigationOptions: ({ navigation }) => ({
       headerTitle: 'New Post',
       headerLeft: (
         <Button
           title="Cancel"
           // Note that since we're going back to a different navigator (CaptureStack -> RootStack)
           // we need to pass `null` as an argument to goBack.
           onPress={() => navigation.goBack(null)}
         />
       ),
     }),
   },
})

export default NewPost

 