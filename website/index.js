let role = 1;
let text = "Ultra Secrete Note Pad"
let index=0;
let title = $("#title");
let speed = 250;
let randomTimeOut = Math.floor(Math.random() * speed);

$(document).ready(function() {

    typeText(text, index);
    changeRole();

});

// SWITCH BETWEEN MESSAGE AND LINK ROLE
const changeRole = () => {

    let submitButton = $("#submitButton");

    submitButton.off("click");
    
    // MESSAGE ROLE
    if (role == 1) {
        
        $(".message").hide();
        $(".link").show();
        console.log('MESSAGE');
        submitButton.click( sendMessage );

        role = 0
    }
    // LINK ROLE
    else {

        $(".message").show();
        $(".link").hide();
        console.log('LINK');
        submitButton.click( sendLink );

        role = 1;
    }
}

$("#toggleRoleButton").click(changeRole);

const typeText = (text, index) => {
    
    if (index < text.length) {
        title.text(title.text() + text.charAt(index));
        
         index++;
         randomTimeOut = Math.floor(Math.random() * speed);

         setTimeout( () => {
            typeText(text, index);

         }, randomTimeOut);
    }
}

// ROLE: MESSAGE -> SEND key-value pair: message-message to server, receive link
function sendMessage() {

    // values from form
    let message = $("#message").val();
    let url = 'http//stronka.pl/';
    let urlkey = '12345678';
    
    // display values to be sent to backend on the console
    console.log(`SENT MESSAGE: ${message}`);
    
    // POST method for now, can be later changed to JSON if needed
    // URL to be updated
    $.post( "http://localhost:3000/go-server-endpoint", {
    
        message: message
        
        }, (response) => {

                // response is JSON object return from the server
                console.log(response);

                // converting JSON into JavaScript object
                let data = JSON.parse(response);

                // display response in console / on website
                // console.log(data.message);
                $("#response").html(`${url} ${urlkey}`);
            }).fail( (jqXHR, textStatus, errorThrown) => {

            console.log('ERROR HADNLING');
            console.log("jqXHR: " + jqXHR);
            console.log("textStatus: " + textStatus);
            console.log("errorThrown: " + errorThrown);

            $("#response").html(`${url}${urlkey}`);
        })
}

// ROLE: MESSAGE -> LINK key-value pair: link-link to server, receive message
function sendLink() {

     // values from form
     let link = $("#link").val();

     let message = 'Siemano';
     
     // display values to be sent to backend on the console
     console.log(`SENT LINK: ${link}`);
     
     // POST method for now, can be later changed to JSON if needed
     // URL to be updated
     $.post( "http://localhost:3000/go-server-endpoint", {
     
        link: link
         
         }, (response) => {
 
                 // response is JSON object return from the server
                 console.log(response);
 
                 // converting JSON into JavaScript object
                 let data = JSON.parse(response);
 
                 // display response in console / on website
                 // console.log(data.message);
                 // $("#response").html(response);
             }).fail( (jqXHR, textStatus, errorThrown) => {
 
             console.log('ERROR HADNLING');
             console.log("jqXHR: " + jqXHR);
             console.log("textStatus: " + textStatus);
             console.log("errorThrown: " + errorThrown);

             $("#response").html(`${message}`);
         })
}
