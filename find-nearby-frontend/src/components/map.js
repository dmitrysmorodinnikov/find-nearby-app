import React from 'react';
import ReactMapGL, { Marker, Popup } from "react-map-gl";

class Map extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      viewport: {
        latitude: this.props.lat !== '' ? Number.parseFloat(this.props.lat) : 1.306002,
        longitude: this.props.lng !== '' ? Number.parseFloat(this.props.lng) : 103.927337,
        width: "100vw",
        height: "100vh",
        zoom: 10
      },
      locations: [],
      selectedVehicle: null
    };
  }

  render() {
    const showCurrentLocationMarker = (this.props.lat !== '' && this.props.lng !== '')
    return (
      <div>
        <ReactMapGL
          {...this.state.viewport}
          mapboxApiAccessToken={process.env.REACT_APP_MAPBOX_TOKEN}
          mapStyle="mapbox://styles/mapbox/dark-v10"
          onViewportChange={vp => {
            this.setState({ viewport: vp })
          }}
        >
          {showCurrentLocationMarker ? (
            <Marker
              latitude={Number.parseFloat(this.props.lat)}
              longitude={Number.parseFloat(this.props.lng)}
            >
              <button className="marker-btn"
              >
                <img src="my-location.svg" alt="current location marker" />
              </button>
            </Marker>
          ) : <div></div>}

          {this.props.locationsToShow.map(loc => (
            <Marker
              key={loc.vehicle_id}
              latitude={loc.latitude}
              longitude={loc.longitude}
            >
              <button className="marker-btn"
                onClick={e => {
                  e.preventDefault();
                  this.setState({ selectedVehicle: loc })
                }}>
                <img src="location-pin.svg" alt="location marker" />
              </button>
            </Marker>
          ))}
          {this.state.selectedVehicle ? (
            <Popup
              latitude={this.state.selectedVehicle.latitude}
              longitude={this.state.selectedVehicle.longitude}
              onClose={() => {
                this.setState({ selectedVehicle: null })
              }}
            >
              <div>
                <h5><span class="badge badge-secondary">Vehicle Info</span></h5>
                <h6><span class="badge badge-light">ID: {this.state.selectedVehicle.vehicle_id}</span></h6>
                <h6><span class="badge badge-light">Lat: {this.state.selectedVehicle.latitude}</span></h6>
                <h6><span class="badge badge-light">Lng: {this.state.selectedVehicle.longitude}</span></h6>
                <h6><span class="badge badge-light">Distance to vehicle: {Math.floor(this.state.selectedVehicle.distance)}m</span></h6>
              </div>
            </Popup>
          ) : null}
        </ReactMapGL>
      </div>
    );
  }
}

export default Map;