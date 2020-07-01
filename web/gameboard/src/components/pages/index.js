import React, { Component } from "react";
import { Button, Container, TextField } from "@material-ui/core";
import { GetCharacters, GetCharacter }from "../middlewares/characters"
import config from "../../config/role-maps.json";
import "./index.css";
class Index extends Component {

    constructor(props) {
        super(props)
        this.state = {
            name: ""
        }
 
    }

    modifyText = (v) => {
        this.setState({
            name: v.target.value
        })
    }

    getCharacters = () => {
        GetCharacters().then(({data}) => {
            console.log(data);
        })
    }

    getCharacter = () => {
        GetCharacter(this.state.name).then(({data}) => {
            window.localStorage.setItem("character", data)
        })
    }

    displayForm = () => {
        return (
            <form noValidate autoComplete="off" className={"form"} onSubmit={this.getCharacter}>
                <TextField 
                    onChange={this.modifyText}
                    placeholder="Name"
                    id="standard-basic" 
                    width="400" 
                    variant="outlined"
                /><br />
                <div className="button-container">
                    <Button onClick={this.getCharacter} variant="contained" color="secondary">Get Role</Button> 
                </div>
            </form>            
        )
    }
    displayCharacter = () => {
        let character = window.localStorage.getItem("character");
        if (character) {
            // display character picture here.
            return (
                <div>
                    <img className="character-portrait" src={`/characters/${character}.jpg`} />
                    <h1>{config[character].text}</h1>
                </div>
            )
        }
        return this.displayForm()
        // display form
    }

    render () {
        return (
            <Container>
                {this.displayCharacter()}
            </Container>            
        )
    }
}

export default Index;