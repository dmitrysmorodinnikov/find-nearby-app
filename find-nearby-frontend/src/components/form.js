import React from 'react';
import axios from 'axios';

class Form extends React.Component {
  constructor(props) {
    super(props)
    this.handleLatChange = this.handleLatChange.bind(this)
    this.handleLngChange = this.handleLngChange.bind(this)
    this.handleRadiusChange = this.handleRadiusChange.bind(this)
    this.handleLimitChange = this.handleLimitChange.bind(this)
    this.handleOnClick = this.handleOnClick.bind(this)
  }

  handleLatChange = (event) => this.props.handleLatChange(event.target.value)
  handleLngChange = (event) => this.props.handleLngChange(event.target.value)
  handleRadiusChange = (event) => this.props.handleRadiusChange(event.target.value)
  handleLimitChange = (event) => this.props.handleLimitChange(event.target.value)


  handleOnClick(event) {
    event.preventDefault();
    axios.get("http://localhost:8081/locations/find", {
      params: {
        latitude: Number.parseFloat(this.props.lat),
        longitude: Number.parseFloat(this.props.lng),
        radius: Number.parseInt(this.props.radius),
        limit: Number.parseInt(this.props.limit)
      }
    }
    ).then(res => {
      this.props.handleLocationsChange(res.data.data)
    })
  }

  render() {
    return (
      <div className="input-group">
        <input id="latitude" name="latitude" type='text' className="form-control text-center" placeholder="latitude* (must be between -/+90)" value={this.props.lat} onChange={this.handleLatChange} />
        <input id="longitude" name="longitude" type='text' className="form-control text-center" placeholder="longitude* (must be between -/+180)" value={this.props.lng} onChange={this.handleLngChange} />
        <input id="radius" name="radius" type='text' className="form-control text-center" placeholder="radius* (must be >= 0)" value={this.props.radius} onChange={this.handleRadiusChange} />
        <input id="limit" name="limit" type='text' className="form-control text-center" placeholder="limit* (must be >= 0)" value={this.props.limit} onChange={this.handleLimitChange} />

        <button className="btn btn-outline-success btn-lg" onClick={this.handleOnClick}>find vehicles</button>
      </div>
    );
  }
}

export default Form;