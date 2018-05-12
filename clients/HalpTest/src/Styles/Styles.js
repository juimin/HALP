// This file should contain the styling required for the different components

const activeTintColor = '#F44336';
const inactiveTintColor = 'gray';

const styles = StyleSheet.create({
   tabBar: {
      height: 49,
      flexDirection: 'row',
      borderTopWidth: StyleSheet.hairlineWidth,
      borderTopColor: 'rgba(0, 0, 0, .4)',
      backgroundColor: '#FFFFFF',
   },
   tab: {
      flex: 1,
      alignItems: 'center',
      justifyContent: 'center',
   },
});