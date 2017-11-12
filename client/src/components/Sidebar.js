import React from 'react'
import PropTypes from 'prop-types'
import { Nav } from 'glamorous'
import { Navlink, NavContainer, NavItem } from '../theme/globalStyle'

const languages = ['Javascript', 'PHP', 'Go', 'Scala', 'Python']
const labels = [
  'api',
  'machine-learning',
  'android',
  'ios',
  'devops',
  'shit-stuff',
]

const Sidebar = ({ onLanguageChange, onLabelChange }) => (
  <Nav
    className="col-sm-3 col-md-2 d-none d-sm-block bg-light"
    position="fixed"
    top="51px"
    bottom={0}
    left={0}
    zIndex={1000}
    padding="20px"
    overflowX="hidden"
    overflowY="auto"
    borderRight="1px solid #eee"
  >
    <NavContainer>
      <NavItem>
        <Navlink>
          Home <span className="sr-only">(current)</span>
        </Navlink>
      </NavItem>
    </NavContainer>
    <NavContainer>
      {languages.map((lang, index) => (
        <NavItem key={index}>
          <Navlink onClick={onLanguageChange(lang)}>{lang}</Navlink>
        </NavItem>
      ))}
    </NavContainer>
    <NavContainer>
      {labels.map((label, index) => (
        <NavItem key={index}>
          <Navlink onClick={onLabelChange(label)}>#{label}</Navlink>
        </NavItem>
      ))}
    </NavContainer>
  </Nav>
)

Sidebar.propTypes = {
  onLanguageChange: PropTypes.func.isRequired,
  onLabelChange: PropTypes.func.isRequired,
}

export default Sidebar
