# resepmasak_be

~ Api Register
    #Body#{
        _"nama": "masukan nama"_,
        _"username": "masukan username"_,
        _"password": "masukan password"_,
        _"email": "masukan email"_,
    }


~ Api Login
    #Body#{
        _"username": "masukan username"_,
        _"password": "masukan password"_
    }

~ API Create Receipt <br>
    "http://127.0.0.1:3000/receipe"<br>
    {
    "title": "Nasi Goreng Bangladesh",<br>
    "description": "Masakan Khas Indonesia dengan cita rasa yang special",<br>
    "ingredients": "siapkan bahan bahan masak, masukan minyak, blendir semua bahan bahan, masukan telor ke wajan yang sudah panas, masukan bumbu yang sudah dihaluskan, oseng oseng sampai bau nya harum, kemudia masukkan nasi yang sudah disiapkan, aduk merata dengan bumbu yang tadi dimasukkan, tambahkan kecap, sajikan",<br>
    "img": ""
}

~ API GET ALL RECEIPE <br>
    http://127.0.0.1:3000/receipe<br>
    Headers: {<br>
        key:LOGIN<br>
        Value: Token<br>
    }<br>


~ API GET BY ID RECIPE <br>
http://127.0.0.1:3000/receipe/get?recipe_id=?<br>

~ API DELETE BY ID RECIPE <br>
http://127.0.0.1:3000/receipe/delete?recipe_id=1 <br>

{
    "code": 200,
    "message": "Receipt berhasil dihapus",
    "status": "success",
    "success": true
}

~ API UPDATE
http://127.0.0.1:3000/receipe/update?recipe_id=1

BODY {
    "title": "Nasi Goreng Banglades",
    "description": "Masakan Khas Indonesia dengan cita rasa yang special",
    "ingredients": "siapkan bahan bahan masak, masukan minyak, blendir semua bahan bahan, masukan telor ke wajan yang sudah panas, masukan bumbu yang sudah dihaluskan, oseng oseng sampai bau nya harum, kemudia masukkan nasi yang sudah disiapkan, aduk merata dengan bumbu yang tadi dimasukkan, tambahkan kecap, sajikan",
    "img": ""
}