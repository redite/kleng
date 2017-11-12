import React from 'react'
import PropTypes from 'prop-types'
import Home from './Home'

const Switcher = ({ currentPage, ...rest }) => {
  let component = <Home {...rest} />

  if (currentPage === 'home') {
    component = <Home {...rest} />
  }

  return (
    <div className="col-sm-9 ml-sm-auto col-md-10 pt-3" role="main">
      {component}
    </div>
  )
}

Switcher.propTypes = {
  currentPage: PropTypes.string.isRequired,
}

Switcher.defaultProps = {
  currentPage: 'home',
}

export default Switcher
