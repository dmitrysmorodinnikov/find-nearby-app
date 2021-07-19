import React from 'react';
import Map from './map';
import Form from './form';

class LocationFinder extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            lat: '',
            lng: '',
            radius: '',
            limit: '',
            locations: []
        }
        this.handleLatChange = this.handleLatChange.bind(this)
        this.handleLngChange = this.handleLngChange.bind(this)
        this.handleLocationsChange = this.handleLocationsChange.bind(this)
    }

    handleLatChange = (changedLat) => this.setState({ lat: changedLat !== '' ? Number.parseFloat(changedLat) : '' })
    handleLngChange = (changedLng) => this.setState({ lng: changedLng !== '' ? Number.parseFloat(changedLng) : '' })
    handleRadiusChange = (changedRadius) => this.setState({ radius: changedRadius !== '' ? Number.parseInt(changedRadius) : '' })
    handleLimitChange = (changedLimit) => this.setState({ limit: changedLimit !== '' ? Number.parseInt(changedLimit) : '' })
    handleLocationsChange = (changedLocations) => this.setState({ locations: changedLocations })

    render() {
        return (
            <div>
                <Form
                    lat={this.state.lat}
                    lng={this.state.lng}
                    radius={this.state.radius}
                    limit={this.state.limit}
                    handleLatChange={this.handleLatChange}
                    handleLngChange={this.handleLngChange}
                    handleRadiusChange={this.handleRadiusChange}
                    handleLimitChange={this.handleLimitChange}
                    handleLocationsChange={this.handleLocationsChange}
                />
                <Map lat={this.state.lat} lng={this.state.lng} locationsToShow={this.state.locations} />
            </div>
        );
    }
}

export default LocationFinder;