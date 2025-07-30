package dsa.searching;

import org.junit.jupiter.api.Test;

import com.dsa.searching.FruitSetExample;

import java.util.HashSet;

import static org.junit.jupiter.api.Assertions.*;

public class HashsetTest {

    @Test
    public void testCreateFruitSet() {
        HashSet<String> fruits = FruitSetExample.getFruitSet();

        assertTrue(fruits.contains("Banana"));
        assertTrue(fruits.contains("Mango"));
        assertFalse(fruits.contains("Apple")); // Apple was removed
        assertEquals(2, fruits.size());
    }
}
