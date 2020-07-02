import React, { Component } from "react";
import { Button, Container, TextField } from "@material-ui/core";
import { GetCharacters, GetCharacter, GetCharacterDetail }from "../middlewares/characters"
import config from "../../config/role-maps.json";
import "./index.css";
class Index extends Component {

    constructor(props) {
        super(props)
        this.state = {
            name: "",
            isLoaded: false
        }
 
    }

    modifyText = (v) => {
        this.setState({
            name: v.target.value
        })
    }

    componentWillMount() {
        this.GetCharacterDetail()
    }

    GetCharacterDetail = () => {
        let sessionId = window.localStorage.getItem("sessionId");
        GetCharacterDetail(sessionId).then(({ data }) => {
            console.log(data)
            this.setState({
                isLoaded: true,
                role: data.message
            })
        })
    }

    getCharacters = () => {
        GetCharacters().then(({data}) => {
            console.log(data);
        })
    }

    getCharacter = () => {
        GetCharacter(this.state.name).then(({data}) => {
            window.localStorage.setItem("sessionId", data.sessId)
            this.setState({
                isLoaded: true,
                role: data.character
            })
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
        if (this.state.isLoaded) {
            if (this.state.role) {
                // display character picture here.
                return (
                    <div>
                        <img className="character-portrait" src={`/characters/${this.state.role}.jpg`} />
                        <h1>{config[this.state.role].text}</h1>
                    </div>
                )
            }
            return this.displayForm()
        }
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