package com.github.ryneal.chaosarcade;

import io.kubernetes.client.openapi.ApiClient;
import org.junit.jupiter.api.Test;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.mock.mockito.MockBean;

@SpringBootTest
class ChaosArcadeApplicationTests {

	@MockBean
	public ApiClient kubernetesClient;

	@Test
	void contextLoads() {
	}

}
