import React from 'react'
import glamorous from 'glamorous'

export const H1 = glamorous.h1({
  marginBottom: '20px',
  paddingBottom: '9px',
  borderBottom: '1px solid #eee',
})

export const Row = glamorous.div('row')

const NavContainerRaw = glamorous.ul('nav nav-pills flex-column')

export const NavContainer = props => (
  <NavContainerRaw {...props} css={{ marginBottom: '20px' }} />
)

const NavItemRaw = glamorous.li('nav-item')

export const NavItem = props => (
  <NavItemRaw {...props} css={{ width: '100%' }} />
)

export const Navlink = glamorous.a({
  borderRadius: 0,
})
