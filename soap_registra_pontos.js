import http from 'k6/http';
import { check } from 'k6';

export const options = {
  vus: 100,
  duration: '30s',
};

const payload = `
<?xml version="1.0" encoding="UTF-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
  <soapenv:Body>
    <registraPontos xmlns="urn:PREPS">
      <autenticacao>teste</autenticacao>
      <codigoEmbarcacao>12345SC</codigoEmbarcacao>
      <pontosEmbarcacao>
        <Ponto>
          <latitude>-27.5969</latitude>
          <longitude>-48.5495</longitude>
          <dataHora>2026-01-20T12:34:56Z</dataHora>
          <leituraSensores>
            <LeituraSensor>
              <tipoSensor>1</tipoSensor>
              <valor>12.3</valor>
            </LeituraSensor>
          </leituraSensores>
        </Ponto>
      </pontosEmbarcacao>
    </registraPontos>
  </soapenv:Body>
</soapenv:Envelope>
`;

export default function () {
  const res = http.post(
    'http://127.0.0.1:8000/rastro',
    payload,
    {
      headers: {
        'Content-Type': 'text/xml; charset=utf-8',
        'SOAPAction': 'urn:RastroAction',
      },
    }
  );

  check(res, {
    'status 200': r => r.status === 200,
  });
}
