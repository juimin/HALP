class TabBar extends Component {
   renderItem = (route, index) => {
     const {
       navigation,
       jumpToIndex,
     } = this.props;
 
     const isNewPost = route.routeName === 'NewPost';
 
     const focused = index === navigation.state.index;
     const color = focused ? activeTintColor : inactiveTintColor;
     const size = 30;
     if (route.routeName == "HomeNav") {
       return (
         <TouchableWithoutFeedback
           key={route.key}
           style={styles.tab}
           onPress={() => jumpToIndex(index)}
         >
           <View style={styles.tab}>
             <Icon style={{ color }} size={size} name="home"/>
           </View>
         </TouchableWithoutFeedback>
       );
     }
     if (route.routeName == "SearchNav") {
       return (
         <TouchableWithoutFeedback
           key={route.key}
           style={styles.tab}
           onPress={() => jumpToIndex(index)}
         >
           <View style={styles.tab}>
             <Icon style={{ color }} size={size} name="search"/>
           </View>
         </TouchableWithoutFeedback>
       );
     }
     if (route.routeName == "AccNav") {
       return (
         <TouchableWithoutFeedback
           key={route.key}
           style={styles.tab}
           onPress={() => jumpToIndex(index)}
         >
           <View style={styles.tab}>
             <Icon style={{ color }} size={size} name="person"/>
           </View>
         </TouchableWithoutFeedback>
       );
     }
     if (route.routeName == "SettingsNav") {
       return (
         <TouchableWithoutFeedback
           key={route.key}
           style={styles.tab}
           onPress={() => jumpToIndex(index)}
         >
           <View style={styles.tab}>
             <Icon style={{ color }} size={size} name="settings"/>
           </View>
         </TouchableWithoutFeedback>
       );
     }
     return (
         <TouchableWithoutFeedback
           key={route.key}
           style={styles.tab}
           onPress={() => isNewPost ? navigation.navigate('NewPostModal') : jumpToIndex(index)}
         >
           <View style={styles.tab}>
             <Icon style={{ color }} style={styles.ico} size={size + 10} name="add-circle"/>
           </View>
         </TouchableWithoutFeedback>
       );
   };
 
   render() {
     const {
       navigation,
     } = this.props;
 
     const {
       routes,
     } = navigation.state;
 
     return (
       <View style={styles.tabBar}>
         {routes && routes.map(this.renderItem)}
       </View>
     );
   }
 }

const Tabs = TabNavigator({
   HomeNav: {
     screen: (props) => <HomeNav />,
   },
   SearchNav: {
     screen: (props) => <SearchNav />,
   },
   NewPost: {
     screen: View,
   },
   AccNav: {
     screen: (props) => <AccNav />,
   },
   SettingsNav: {
     screen: (props) => <SettingsNav />,
   },
 }, {
   tabBarPosition: 'bottom',
   tabBarComponent: TabBar,
 });

export default Tabs;