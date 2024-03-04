<?php

$curl = curl_init();

curl_setopt_array($curl, array(
  CURLOPT_URL => 'https://api-key.fusionbrain.ai/key/api/v1/text2image/run',
  CURLOPT_RETURNTRANSFER => true,
  CURLOPT_ENCODING => '',
  CURLOPT_MAXREDIRS => 10,
  CURLOPT_TIMEOUT => 0,
  CURLOPT_FOLLOWLOCATION => true,
  CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
  CURLOPT_CUSTOMREQUEST => 'POST',
  CURLOPT_POSTFIELDS => array('model_id' => '4','params'=> new CURLFILE('WNeLCvmzE/generate.json')),
  CURLOPT_HTTPHEADER => array(
    'X-Key: Key FFDB0757E5E5A2FF2D7A297CE95BDA2D',
    'X-Secret: Secret F9C4980BC3166E19C9A42607358B8DA6',
    'X-API-Key: {{token}}'
  ),
));

$response = curl_exec($curl);

curl_close($curl);
echo $response;
