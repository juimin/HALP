import React from 'react';
import PropTypes from 'prop-types';
import {
  View,
} from 'react-native';

const HideableView = (props) => {
  const { children, hide, style } = props;
  if (hide) {
    return null;
  }  
  if (children) {
    return (
        <View {...this.props} style={style}>
          { children }
        </View>
      );
  } else {
    return (
        <View {...this.props} style={style} />
    );
    }
};

HideableView.propTypes = {
  children: PropTypes.oneOfType([
    PropTypes.string,
    PropTypes.element,
    PropTypes.number,
    PropTypes.arrayOf(PropTypes.oneOfType([
      PropTypes.string,
      PropTypes.number,
      PropTypes.element,
    ])),
  ]).isRequired,
  style: View.propTypes.style,
  hide: PropTypes.bool,
};

export default HideableView;