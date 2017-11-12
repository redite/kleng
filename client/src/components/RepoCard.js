import React from 'react'
import PropTypes from 'prop-types'
import { Div, H4, P, A } from 'glamorous'

const RepoCard = ({ repo }) => {
  return (
    <Div className="col">
      <Div className="card p-2" css={{ width: '20rem' }}>
        <Div className="card-body">
          <H4 className="card-title">{repo.name}</H4>
          <P>{repo.description}</P>
          <A href={repo.url} className="btn btn-primary">
            Go to github
          </A>
        </Div>
      </Div>
    </Div>
  )
}

RepoCard.propTypes = {
  repo: PropTypes.shape({
    name: PropTypes.string,
    username: PropTypes.string,
    description: PropTypes.string,
    url: PropTypes.string,
  }).isRequired,
}

export default RepoCard
