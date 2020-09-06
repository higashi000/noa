function MakeChannel() {
   let channelName = document.getElementById('name').value;
   let adminName = document.getElementById('admin').value;
   let fileType = document.getElementById('filetype').value;

   let data = {
      roomid: channelName,
      admin: adminName,
      filetype: fileType,
   };

   fetch("/noa/registchannel", {
      method: 'POST',
      body: JSON.stringify(data),
      headers: {
         'Content-Type': 'application/json'
      },
      mode: 'no-cors'
   })
      .then(res => res.json())
      .catch(error => console.eror('Error:', error));
}
