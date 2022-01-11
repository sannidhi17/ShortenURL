import React from "react";
import './URLInput.css';

class URLInput extends React.Component {
    constructor (props) {
        super(props);
        this.state = {
            url: '', 
            shorturl: ''
        }

        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleChange(event) {
        this.setState({
            url: event.target.value,
            shorturl: this.state.shorturl
        });
    }

    handleSubmit (event) {
        console.log(this.state.url)
        event.preventDefault();

        fetch("http://localhost:8080/create-short-url", {
            "method": "POST",
            "headers": {
                "Content-Type": "application/json",
                "accept": "application/json"
            },
            "body": JSON.stringify({"original_url": this.state.url})
        })
        .then(response => response.json())
        .then(response => {
            console.log(response.short_url)
            this.setState({
                url: this.state.url,
                shorturl: response.short_url
            })
        })
        .catch(err => {
            console.log(err)
        })
    }

    render() {
        return (
            <div className="FullWidth">
                <form onSubmit={this.handleSubmit} className="FullWidth">
                    <input type="text" value={this.state.value} onChange={this.handleChange} placeholder="https://sample-url.com" className="TextInput"/>
                    <input type="submit" value="SUBMIT" className="SubmitButton"/>
                </form>
                <h4>
                    <a href = {this.state.shorturl} target="_blank" className="App-link">{this.state.shorturl}</a>
                </h4>
            </div>
        );
    }
}

export default URLInput;